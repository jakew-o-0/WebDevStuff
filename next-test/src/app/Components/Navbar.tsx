'use client'

import React from 'react'
import Image from 'next/image'
import Link from 'next/link'
import { UserButton, useUser } from '@clerk/nextjs'


export default function Navbar() {
  const {user, isLoaded} = useUser();

  return (
    <nav className='sticky flex top-0 my-2 mx-auto bg-slate-100 rounded shadow-md lg:max-w-5xl'>
      
        {/* Logo */}
        <Link href={'/'} className='flex-grow flex'>
            <Image
            src={'/logo.png'}
            alt='logo'
            width={80}
            height={80}
            className='p-2'
            />
            <div className='my-auto font-semibold text-xl'>
                <h1>Gibjohn</h1>
                <h2>Tutoring</h2>
            </div>
        </Link>

        {/* Buttons when user is not logged in*/}
        {
          isLoaded === true && !user &&
          <div className='flex my-auto'>
            <Link href={'/sign-in'}>
              <button className='flex items-center p-1 mx-2 rounded text-xl bg-indigo-700 text-gray-100 hover:bg-indigo-600 active:ring active:ring-indigo-600 '>Sign-in</button>
            </Link>
            <Link href={'/sign-up'}>
              <button className='flex items-center p-1 mx-2 rounded text-xl bg-indigo-700 text-gray-100 hover:bg-indigo-600 active:ring active:ring-indigo-600 '>Sign-up</button>
            </Link>
          </div>
        }

        {/* Buttons when user is logged in*/}
        {
          isLoaded === true && user && 
          <>
            <Link href={'/user/home'} className='my-auto'>
              <button className='flex items-center p-1 rounded text-xl bg-indigo-700 text-gray-100 hover:bg-indigo-600 active:ring active:ring-indigo-600 '>Home</button>
            </Link>
            <div className='my-auto mx-2'>
              <UserButton appearance={{elements: {avatarBox: 'w-16 h-auto'}}}/>
            </div>
          </>
        }
    </nav>
  )
}
