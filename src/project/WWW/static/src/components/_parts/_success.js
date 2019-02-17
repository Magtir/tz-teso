import React from 'react'
import './_success.scss'
import imgSuccess from '../../images/success.png'

class _success extends React.Component {
    render() {
        return (
            <div className={'success ' + this.props.status}>
                <div className="wrapper"
                     style={{backgroundImage:'url(../WWW/static/dist/' + imgSuccess + ')'}}>
                </div>
            </div>
        )
    }
}

export default _success