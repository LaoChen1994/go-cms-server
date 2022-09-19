import type { createBrowserRouter } from 'react-router-dom'

type ParamsType = Parameters<typeof createBrowserRouter>

export interface IGetRoutes {
    (): {
        routes: ParamsType[0],
        options?: ParamsType[1]
    }
}
