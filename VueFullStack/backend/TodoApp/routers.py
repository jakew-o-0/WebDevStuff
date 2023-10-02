from fastapi import APIRouter, Body, HTTPException, Request, status
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse

from .models import TaskModel, UpdateTaskModel

router = APIRouter()


@router.post('/', response_description="Add new task")
async def create_task(request: Request, task: TaskModel):
    task = jsonable_encoder(task)
    db = request.app.mongodb

    new_task = await db['Tasks'].insert_one(task)
    created_task = await db['Tasks']\
        .find_one( {'_id': new_task.inserted_id} )

    return JSONResponse(status_code=status.HTTP_201_CREATED, 
                        content=created_task)



@router.get('/', response_description="Get all tasks")
async def get_all_tasks(request: Request):
    db = request.app.mongodb
    tasks = []

    for task in await db['Tasks'].find().to_list(length=100):
        tasks.append(task)
    
    return tasks



@router.get('/{id}', response_description="Get task by id")
async def get_task_id(id: str, request: Request):
    db = request.app.mongodb
    
    if (task := await db['Tasks'].find_one({'_id': id})) is not None:
        return task

    return HTTPException(status_code=status.HTTP_404_NOT_FOUND, 
                            detail=f"task with id: '{id}' does not exist")

    

@router.put('/{id}', response_description="Update a task")
async def update_task(id: str, request: Request, update_task: UpdateTaskModel):
    db = request.app.mongodb

    if update_task.name is not None:
        is_name_updated = await db['Tasks'].update_one({'_id': id}, 
                                                       {'$set': {'name': update_task.name}})

    if update_task.completed is not None:
        is_completed_updated = await db['Tasks'].update_one({'_id': id}, 
                                                            {'$set': {'completed': update_task.completed}})

    if  (is_name_updated.modified_count >= 1 or\
        is_completed_updated.modified_count >= 1) and\
        ((updated_task := await db['Tasks'].find_one({'_id': id})) is not None):
            return JSONResponse(status_code=status.HTTP_200_OK, 
                                content=updated_task) 

    return HTTPException(status_code=status.HTTP_404_NOT_FOUND, 
                         detail=f"task with id: '{id}' does not exist")



@router.delete('/{id}', response_description="remove a task")
async def delete_task(id: str, request: Request):
    db = request.app.mongodb
    deleted_task = await db['Tasks'].delete_one({'_id': id})

    if deleted_task.deleted_count >= 1:
        return HTTPException(status_code=status.HTTP_204_NO_CONTENT,
                             detail="task Successfully deleted")

    return HTTPException(status_code=status.HTTP_404_NOT_FOUND, 
                         detail=f"task with id: '{id}' does not exist")