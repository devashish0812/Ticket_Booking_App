import logo from "../assets/eventflow-logo-removebg.png";
function Dashboard() {
  

  return (
    <>
<div className="flex flex-cols-3 gap-8">
    <div className="w-1/4" ></div>
    <div className="w-1/2 flex justify-center" >
        <img src={logo} className="w-40 h-auto"></img>
    </div>
    <div className="w-1/4" ></div>
</div>

<div className="flex gap-8 p-8">
  <div className="w-1/6 bg-gray-100 p-4 rounded">
    <h2 className="font-bold mb-4">Filters</h2>
    <div>
      <label className="block mb-2">Category</label>
      <select className="w-full p-2 border rounded">
        <option>Music</option>
        <option>Sports</option>
        <option>Workshop</option>
      </select>
    </div>
    <div className="mt-4">
      <label className="block mb-2">Date</label>
      <input type="date" className="w-full p-2 border rounded" />
    </div>
  </div>


  <div className="w-2/3">
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">

      <div className="bg-white shadow p-4 rounded">
        Event 1
      </div>
      <div className="bg-white shadow p-4 rounded">
        Event 2
      </div>
      <div className="bg-white shadow p-4 rounded">
        Event 3
      </div>
    </div>
  </div>

      <div className="w-1/6"></div>

</div>



    </>
  )
}

export default Dashboard
