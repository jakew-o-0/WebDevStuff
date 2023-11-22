import React from 'react'

export default function Button({ children }: { children: React.ReactNode}, colour: string, extraStyling: string,) {
    let style = 'flex items-center p-1 rounded'

    if(colour === 'violet') {
        style.concat('bg-violet-700 text-gray-100 hover:bg-violet-600 active:ring active:ring-violet-600')
    } 
    else if(colour === 'indigo') {
        style.concat(' bg-indigo-700 text-gray-100 hover:bg-indigo-600 active:ring active:ring-indigo-600 ')
    } 
    else if(colour === 'green') {
        style.concat(' bg-emerald-300 text-stone-800 hover:bg-emerald-400 active:ring active:ring-violet-400 ')
    }

    style.concat(' ' + extraStyling)

    return (
        <button className={style}>
            {children}
        </button>
    )
}
