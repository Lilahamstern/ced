import React, { Component } from 'react'
import { BrowserRouter as Router, Link } from 'react-router-dom'

export class Navbar extends Component {
  render() {
    return (
      <Router>
        <header className="flex justify-between items-center bg-gray-800 px-4 py-3">
          <div className="mr-8">
            <span className="text-3xl font-semibold text-white">CED</span>
          </div>

          <div className="text-left w-full hidden block flex-grow lg:flex lg:items-center lg:w-auto">
            <div className="text-base lg:flex-grow">
              <Link to="/" className="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4">Projects</Link>
            </div>
          </div>

          <div className="lg:hidden">
            <button className="text-gray-600 hover:text-white focus:text-white focus:no-underline">
              <svg className="h-8 w-8 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path fillRule="evenodd" d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z" />
              </svg>
            </button>
          </div>
        </header>
      </Router>
    )
  }
}

export default Navbar
