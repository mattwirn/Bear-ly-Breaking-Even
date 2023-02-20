import Image from 'next/image'


function Logo({w}: SafeNumber) {
    return (
      <Image 
      
        src='/logo.png'
        height={0} 
        width={w}
        alt='home'
      />
    )
}

export default Logo