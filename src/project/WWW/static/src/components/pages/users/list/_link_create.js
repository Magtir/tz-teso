import React from 'react'
import {Link} from 'react-router-dom'
import imgCreate from '../../../../images/create.png'

class LinkCreate extends React.Component {
    render() {
        return (
            <Link className={'link-create'}
                  to={'/user/create'}
                  style={{backgroundImage: "url(../WWW/static/dist/" + imgCreate + ")"}}>
            </Link>
        )
    }
}

export default LinkCreate