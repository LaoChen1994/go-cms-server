import { createRoot } from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { AuthProvider } from 'Context/auth'

import routes from "Config/router";
import NotFount from "Pages/common/NotFount";

function App() {
  const router = createBrowserRouter([routes])

  return (
    <AuthProvider>
      <RouterProvider router={router} fallbackElement={<NotFount />} />
    </AuthProvider>
  )
}

document.title = "皮蛋的CMS后台"
createRoot(document.getElementById('app')!).render(<App />);
