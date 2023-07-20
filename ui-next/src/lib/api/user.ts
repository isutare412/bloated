import 'server-only'
import { getApiBaseUrl } from '@/lib/env'
import { User } from '@/lib/model/user'
import { wait } from '@/lib/utils/wait'

export async function getUserById(id: number): Promise<User> {
  const response = await fetch(`${getApiBaseUrl()}/api/v1/users/${id}`, {
    cache: 'no-store',
  })
  if (!response.ok) {
    throw new Error(`failed to get user by id(${id})`)
  }

  await wait(1000)
  return response.json()
}
