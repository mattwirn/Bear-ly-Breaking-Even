import {useRouter} from 'next/router'
import Image from 'next/image'
import Logo from '@/components/Logo'

function PageHeader() {
    const router = useRouter()

    function logLink() {
        router.push('/login')
    }
    
    function signupLink() {
      router.push('/signup')
    }
    
    function homeLink() {
        router.push('/')
    }

    return (
        <div className='pageHeader'>
            <div className='pt-1'>
                <button onClick={homeLink} className='hover:underline mx-1'>
                    <Logo w={50}/>
                </button>
            </div>
            <div className='text-lg'>
                <button onClick={signupLink} className='hover:underline mx-1'>Sign up</button>
                <button onClick={logLink} className='hover:underline mx-1'>Log in</button>
            </div>
        </div>
    )
}

export default PageHeader