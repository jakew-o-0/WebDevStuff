import React from 'react'
import Image from 'next/image'
import StyledButton from './Button'

export default function Navbar() {
  return (
    <>
    <nav className='sticky flex top-0 my-2 mx-auto bg-slate-100 rounded shadow-md lg:max-w-5xl'>
        <div className='flex-grow flex'>
            <Image
            src={'/logo.png'}
            alt='logo'
            width={70}
            height={70}
            className='p-2'
            />
            <div className='my-auto font-semibold text-xl'>
                <h1>Gibjohn</h1>
                <h2>Tutoring</h2>
            </div>
        </div>
        <div>
            <StyledButton colour={'indigo'} extraStyling={'my-auto mx-2'}> login </StyledButton>
        </div>
    </nav>
    </>
  )
}
