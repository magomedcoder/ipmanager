import { createGrpcWebTransport } from '@connectrpc/connect-web'
import { createClient } from '@connectrpc/connect'
import type { Client } from '@connectrpc/connect'

const transport = createGrpcWebTransport({
  baseUrl: import.meta.env.VITE_BASE_API,
  useBinaryFormat: true
})

export const client = <T>(service: T): Client<T> => createClient(service, transport)
