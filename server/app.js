import path from 'path'
import fastify from 'fastify'
import fastifyStatic from 'fastify-static'
import helmet from 'fastify-helmet'
import hyperid from 'hyperid'
import { enableCORS, serveIndex } from './util'
import { init as uploadProviderInit } from './uploads'
import api from './api'

const app = fastify({
  logger: {
    level: process.env.NODE_ENV === 'production' ? 'info' : 'debug'
  },
  genReqId: hyperid()
})

app.register(enableCORS)
app.register(helmet, {
  dnsPrefetchControl: false,
  referrerPolicy: { policy: 'same-origin' },
  contentSecurityPolicy: {
    directives: {
      defaultSrc: ["'self'"],
      fontSrc: ['fonts.gstatic.com', "'self'", 'data:'],
      styleSrc: ['fonts.googleapis.com', "'unsafe-inline'", "'self'"],
      imgSrc: ['*', 'data:']
    }
  }
})

uploadProviderInit(app)

app.register(api, { prefix: '/api/v1/' })

const staticPath = path.join(__dirname, '../build')

app.register(serveIndex, {
  indexPath: path.join(staticPath, 'index.html')
})
app.register(fastifyStatic, {
  root: staticPath
})

export default app
