import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import '@/app/globals.css'
import { ReduxProvider } from '@/components/reduxProvider'
import { ToastContainer } from '@/components/toast'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Redshore',
  manifest: '/site.webmanifest',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <head>
        <meta name="msapplication-TileColor" content="#da532c" />
        <meta name="theme-color" content="#ffffff" />
      </head>
      <body
        className={`flex min-h-screen flex-col justify-start ${inter.className}`}
      >
        <ReduxProvider>
          <div className="mx-auto w-full max-w-3xl p-3">{children}</div>
          <footer className="footer mt-auto bg-base-200 p-8 text-base-content">
            <div className="mx-auto w-full max-w-3xl">
              Copyright &copy; Redshore 2023
            </div>
          </footer>
          <ToastContainer />
        </ReduxProvider>
      </body>
    </html>
  )
}
