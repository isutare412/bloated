import { revalidatePath } from 'next/cache'
import { NextRequest, NextResponse } from 'next/server'

export async function GET(request: NextRequest) {
  revalidatePath('/users')
  return NextResponse.json({ revalidated: true, now: Date.now() })
}
