function RenderButtons({ IsLoggedIn }) {
    if(!IsLoggedIn) {
        return (
            <div className="ml-auto mr-2 my-auto flex justify-center">
                <button className="bg-violet-700 text-gray-100  flex items-center p-1 mx-2 my-2 rounded hover:bg-violet-600 active:ring active:ring-violet-600 ">
                    <p className="text-lg font-Body">
                        Login
                    </p>
                </button>
                <button className="bg-violet-700 text-gray-100  flex items-center p-1 mx-2 my-2 rounded hover:bg-violet-600 active:ring active:ring-violet-600 " >
                    <p className="text-lg font-Body">
                        SignUp
                    </p>
                </button>
            </div>
        );
    }
    return(
        <div className="ml-auto mr-2 my-auto flex justify-center">
            <button className="bg-violet-700 text-gray-100  flex items-center p-1 mx-2 my-2 rounded hover:bg-violet-600 active:ring active:ring-violet-600 " >
                <svg xmlns="http://www.w3.org/2000/svg" height="30" width="30" viewBox="0 0 24 24" fill="currentColor">
                    <path fill-rule="evenodd" d="M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" clip-rule="evenodd" />
                </svg>
                <p className="text-lg font-Body"> Home </p>
            </button >
        </div >
    );
}

function Navbar({ IsLoggedIn }) {
    return (  
        <nav className="sticky top-0 flex w-auto h-full font-Header m-5 rounded-md shadow-md bg-gray-100">

            <div className="flex hover:bg-gray-200 w-fit rounded">
                <svg className="m-2" width="75" height="75" viewBox="0 0 300 300" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <g clip-path="url(#clip0_13_3)">
                        <path opacity="0.7" fill-rule="evenodd" clip-rule="evenodd" d="M50.0398 255.304C19.0288 227.671 1.44849 190.292 1.08775 151.223C0.72702 112.153 17.6147 74.522 48.1113 46.4385C48.9791 45.6655 50.2885 45.6655 51.1564 46.4384L281.887 251.918C284.125 253.911 284.122 257.417 281.809 259.323C250.888 284.79 210.244 298.976 167.983 298.911C123.74 298.843 81.3305 283.163 50.0398 255.304Z" fill="#4B00FF" />
                        <path opacity="0.8" fill-rule="evenodd" clip-rule="evenodd" d="M50.04 43.6061C81.0642 15.9849 123.031 0.326346 166.895 0.00504352C210.759 -0.316259 253.008 14.7254 284.538 41.8884C285.349 42.6106 285.349 43.8784 284.539 44.6007L52.9718 250.884C51.071 252.577 48.1944 252.578 46.3483 250.825C17.2395 223.189 1.00792 186.657 1.0813 148.657C1.1574 109.251 18.7619 71.4764 50.04 43.6061Z" fill="#F93535" />
                        <path fill-rule="evenodd" clip-rule="evenodd" d="M49.6338 253.878C49.6338 253.865 49.6393 253.852 49.6488 253.844L162.67 153.206C164.904 151.217 164.903 147.724 162.668 145.736L51.1572 46.558C50.2887 45.7855 48.9793 45.786 48.1113 46.5591C17.7028 74.3697 0.813241 111.727 1.09818 150.547C1.38289 189.334 18.7913 226.463 49.5587 253.911C49.5877 253.937 49.6338 253.916 49.6338 253.878Z" fill="#4B00FF" />
                        <path opacity="0.8" d="M212.033 131.236C212.033 128.474 214.272 126.236 217.033 126.236H295C297.761 126.236 300 128.474 300 131.236V293.005C300 295.766 297.761 298.005 295 298.005H217.033C214.272 298.005 212.033 295.766 212.033 293.005V131.236Z" fill="#41EFB5" />
                    </g>
                    <defs>
                        <clipPath id="clip0_13_3">
                            <rect width="300" height="300" fill="white" />
                        </clipPath>
                    </defs>
                </svg>

                <div className="my-auto">
                    <h1 className="text-4xl font-bold text-stone-800">Gibjohn</h1>
                    <h2 className="mx-1 text-xl font-semibold text-stone-800">Tutoring</h2>
                </div>
            </div>

            <RenderButtons IsLoggedIn={IsLoggedIn}/>

        </nav >
    );
}

export default Navbar;