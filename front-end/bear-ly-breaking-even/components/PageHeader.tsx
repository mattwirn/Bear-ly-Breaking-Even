import {useRouter} from 'next/router'
import Logo from '@/components/Logo'

function PageHeader({display}: IntrinsitAttributes) {
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

    function dashLink() {
        router.push('/dashboard')
    }

    return (
        <div className='pageHeader'>
            <div className='pt-1'>
                <button onClick={homeLink} className='hover:underline'>
                    <Logo w={90}/>
                </button>
            </div>
            { display ? 
            <div className='text-lg'>
                <button onClick={signupLink} className='hover:underline mx-1'>Sign up</button>
                <button onClick={logLink} className='hover:underline mx-1'>Log in</button>
            </div>
            : <div></div> }
        </div>
    )
}

export default PageHeader