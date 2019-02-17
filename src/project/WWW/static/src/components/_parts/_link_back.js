import React from 'react'
import {Link} from 'react-router-dom'
import './_link_back.scss'

class LinkBack extends React.Component {
    render() {
        return (
            <Link className={'link-back'} to={'/users'}>
                Назад
            </Link>
        )
    }
}

export default LinkBack