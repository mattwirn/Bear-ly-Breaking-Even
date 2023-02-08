import {useRouter} from 'next/router'

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
            <div className=''>
                <button onClick={homeLink} className='hover:underline mx-1'>Home</button>
            </div>
            <div>
                <button onClick={signupLink} className='hover:underline mx-1'>Sign up</button>
                <button onClick={logLink} className='hover:underline mx-1'>Log in</button>
            </div>
        </div>
    )
}

export default PageHeader