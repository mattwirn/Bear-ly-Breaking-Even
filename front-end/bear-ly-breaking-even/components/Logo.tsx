import Image from 'next/image'


function Logo({w}: SafeNumber, {h}: SafeNumber) {
    return (
      <img 
        data-cy='logo'
        src='/logo.png'
        height={h} 
        width={w}
        alt='home'
      />
    )
}

export default Logo