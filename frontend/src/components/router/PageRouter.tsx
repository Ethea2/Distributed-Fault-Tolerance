import { Route, Routes } from "react-router";
import routes from "./router";

export default function PageRouter() {
  return (
    <>
      <Routes>
        {
          routes.map((route, index) => (
            <Route
              path={route.path}
              element={route.element}
              key={index}
            />
          ))
        }
      </Routes>
    </>
  )
}
