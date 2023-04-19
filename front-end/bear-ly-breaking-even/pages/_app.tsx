import '@/styles/globals.css'
import type { AppProps } from 'next/app'

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <div className='min-h-screen min-w-min min-w-screen bg-[#deb887]'>
      <Component {...pageProps} />
    </div>
  )
}
