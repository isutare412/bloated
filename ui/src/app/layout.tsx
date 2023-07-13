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
      <body className={inter.className}>
        <ReduxProvider>
          <div className="mx-auto max-w-3xl p-3">{children}</div>
          <ToastContainer />
        </ReduxProvider>
      </body>
    </html>
  )
}
