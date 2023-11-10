function HomePage() {
    return (  
        <div className="flex justify-center m-auto">
            <div className="flex-row mr-5 justify-center">
                <div className="mx-auto p-5 bg-emerald-300 rounded shadow-md h-fit">
                    <h1 className=" font-semibold text-5xl font-Header">Collaborative learning platform</h1>
                    <ul className="mt-5 list-disc list-inside font-Body">
                        <li>Work togather on projects</li>
                        <li>Compete with your class to earn points for your group</li>
                    </ul>
                </div>

                <div className="mx-auto my-5 p-5 bg-red-400 rounded shadow-md h-fit w-fit">
                    <p className=" font-Body">login now and start collaborating</p>
                </div>
            </div>

            <img className="rounded" src="../assets/teacher-7408666.svg" alt="" width="500"/>
        </div>
    );
}

export default HomePage;