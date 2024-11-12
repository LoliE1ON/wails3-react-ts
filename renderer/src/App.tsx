import {WindowPanel} from "./components/base/WindowPanel.tsx";

export const App = () => {
  return (
      <>
          <WindowPanel/>

          <div className="flex h-screen">
              <div className="self-start mt-24 ml-8">Page</div>
          </div>
      </>
  )
}
