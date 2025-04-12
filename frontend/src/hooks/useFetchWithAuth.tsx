import { useEffect, useState } from "react"
import { toast } from "react-toastify"

const useFetchWithAuth = (url: string) => {
  const [data, setData] = useState<any>()
  const [loading, setLoading] = useState<boolean>(false)

  const refetch = async () => {
    setLoading(true)
    try {
      const res = await fetch(`${import.meta.env.VITE_API_URL}${url}`, {
        headers: {
          "Authorization": `Bearer ${JSON.parse(localStorage.getItem("user")!).token}`
        }
      })

      const json = await res.json()

      if (!res.ok) {
        console.log(json)
        toast("Something went wrong!", {
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: false,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "dark",
          type: "error"
        })
      }
      setData(json)
      setLoading(false)
    } catch (e) {
      setLoading(false)
      const err = e as Error
      console.log(err.message)
      toast("Something went wrong!", {
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: false,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
        type: "error"
      })
    }
  }


  useEffect(() => {
    const fetchData = async () => {
      setLoading(true)
      try {
        const res = await fetch(`${import.meta.env.VITE_API_URL}${url}`, {
          headers: {
            "Authorization": `Bearer ${JSON.parse(localStorage.getItem("user")!).token}`
          }
        })

        const json = await res.json()

        if (!res.ok) {
          console.log(json)
          toast("Something went wrong!", {
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: false,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "dark",
            type: "error"
          })
        }
        setData(json)
        setLoading(false)
      } catch (e) {
        setLoading(false)
        const err = e as Error
        console.log(err.message)
        toast("Something went wrong!", {
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: false,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "dark",
          type: "error"
        })
      }
    }
    fetchData()
  }, [url])

  return { data, loading, refetch }
}

export default useFetchWithAuth 
