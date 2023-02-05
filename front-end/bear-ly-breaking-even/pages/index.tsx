import Head from 'next/head'
import styles from '@/styles/Home.module.css'


export default function Home() {
  return (
    <div className=''>
      <Head>
        <title>Bear-ly Breaking Even</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport"/>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className='flex w-screen h-screen items-center justify-center'>
        <button className='bg-[#addfad] rounded-lg pl-5 pr-5 pt-2 pb-2 border border-black'>
          <div className=''>Login</div>
        </button>
        <div>

        </div>
      </div>
    </div>
  )
}
