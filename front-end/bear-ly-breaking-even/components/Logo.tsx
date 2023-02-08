import Image from 'next/image'


function Logo() {
    return (
        <Image 
        src='/logo.png'
        height={30} 
        width={150}
        alt='home'
      />
    )
}

export default Logo