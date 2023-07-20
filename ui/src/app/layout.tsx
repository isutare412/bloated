import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import '@/app/globals.css'
import { Provider } from '@/components/reduxProvider'
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
      <body className={`flex min-h-screen flex-col ${inter.className}`}>
        <Provider>
          <div className="mx-auto w-full max-w-3xl flex-1 p-3">{children}</div>
          <footer className="footer flex-none bg-base-200 p-6 text-base-content">
            <div className="mx-auto w-full max-w-3xl">
              Copyright &copy; Redshore 2023
            </div>
          </footer>
          <ToastContainer />
        </Provider>
      </body>
    </html>
  )
}
