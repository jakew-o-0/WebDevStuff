from fastapi import APIRouter, Body, HTTPException, Request, status, Depends
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
from typing import Annotated

from .models import TaskModel, UpdateTaskModel

router = APIRouter()


@router.post('/', response_description="Add new task")
async def create_task(request: Request, task: TaskModel):
    task = jsonable_encoder(task)
    cluster = request.app.mongodb['Tasks']

    new_task = await cluster.insert_one(task)
    created_task = await cluster.find_one( {'_id': new_task.inserted_id} )

    return JSONResponse(
        status_code=status.HTTP_201_CREATED, 
        content=created_task
    )



@router.get('/', response_description="Get all tasks")
async def get_all_tasks(request: Request):
    cluster = request.app.mongodb['Tasks']
    tasks = []

    for task in await cluster.find().to_list(length=100):
        tasks.append(task)
    
    if tasks == [] :
        return HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Tasks Empty",
        )
    else:
        return tasks



@router.get('/{id}', response_description="Get task by id")
async def get_task_id(id: str, request: Request):
    cluster = request.app.mongodb['Tasks']
    
    if (task := await cluster.find_one({'_id': id})) is not None:
        return task

    return HTTPException(
        status_code=status.HTTP_404_NOT_FOUND, 
        detail=f"task with id: '{id}' does not exist"
    )

    

@router.put('/{id}', response_description="Update a task")
async def update_task(id: str, request: Request, update_task: UpdateTaskModel):
    cluster = request.app.mongodb['Tasks']
    tasks_updated = 0

    if update_task.name is not None:
        is_name_updated = await cluster.update_one(
            {'_id': id}, 
            {'$set': {'name': update_task.name}}
        )
        if is_name_updated.modified_cound >= 1:
            tasks_updated += 1

    if update_task.completed is not None:
        is_completed_updated = await cluster.update_one(
            {'_id': id}, 
            {'$set': {'completed': update_task.completed}}
        )
        if is_completed_updated.modified_count >= 1:
            tasks_updated += 1

    if  (tasks_updated >= 1):
        return HTTPException(
            status_code=status.HTTP_200_OK,
            detail="task updated"
        )

    return HTTPException(
        status_code=status.HTTP_404_NOT_FOUND, 
        detail=f"task with id: '{id}' does not exist"
    )



@router.delete('/{id}', response_description="remove a task")
async def delete_task(id: str, request: Request):
    cluster = request.app.mongodb['Tasks']
    deleted_task = await cluster.delete_one({'_id': id})

    if deleted_task.deleted_count >= 1:
        return HTTPException(
            status_code=status.HTTP_204_NO_CONTENT,
            detail="task Successfully deleted"
        )

    return HTTPException(
        status_code=status.HTTP_404_NOT_FOUND, 
        detail=f"task with id: '{id}' does not exist"
    )