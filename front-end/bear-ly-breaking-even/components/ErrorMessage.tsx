import React from 'react'

export default function ErrorMessage({taken}: IntrinsitAttributes) {
    
    if (taken) {
    return (
    <div id='errorMessage' className='mt-3 rounded w-52 border border-red-300'>
        <div className='p-2 text-red-500 font-bold font-mono text-sm'>
          The username you entered is taken. Please choose another.
        </div>
    </div>
    )
    }
    else {
        return <div></div>
    }
}
