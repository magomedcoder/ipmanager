import { createGrpcWebTransport } from '@connectrpc/connect-web'
import { Code, createClient } from '@connectrpc/connect'
import type { Client } from '@connectrpc/connect'
import type { Interceptor } from '@connectrpc/connect'
import { ConnectError } from '@connectrpc/connect'
import { getAccessToken, removeAccessToken } from '@/utils/cookie'

const auth: Interceptor = (next) => async (req) => {
  try {
    const token = getAccessToken()
    if (token) {
      req.header.append('Authorization', `Bearer ${token}`)
    }

    return await next(req)
  } catch (err) {
    if (err instanceof ConnectError) {
      if (err.code == Code.Unauthenticated) {
        removeAccessToken()
        location.reload()
      }
    }
    throw err
  }
}

const transport = createGrpcWebTransport({
  baseUrl: import.meta.env.VITE_BASE_API,
  useBinaryFormat: true,
  interceptors: [auth]
})

export const client = <T>(service: T): Client<T> => createClient(service, transport)
