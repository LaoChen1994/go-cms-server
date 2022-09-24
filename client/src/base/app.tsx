import { createRoot } from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import routes from "Config/router";

function App() {
  const router = createBrowserRouter([routes])

  return (
    <RouterProvider router={router} />
  )
}

document.title = "皮蛋的CMS后台"
createRoot(document.getElementById('app')!).render(<App />);
