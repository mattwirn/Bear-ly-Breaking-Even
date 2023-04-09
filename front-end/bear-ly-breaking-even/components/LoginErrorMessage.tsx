import React from 'react'

export default function LoginErrorMessage({exists}: IntrinsitAttributes) {
    if (!exists) {
        return (
        <div id='errorMessage' className='mt-3 rounded w-52 border border-red-300'>
            <div className='p-2 text-red-500 font-bold font-mono text-sm'>
              Invalid username or password.
            </div>
        </div>
        )
    }
    else {
        return <div></div>
    }
}
