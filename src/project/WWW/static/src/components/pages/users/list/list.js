import React from "react";
import RequestServer from '../../../lib/request_server'
import {Link} from 'react-router-dom'
import "./list.scss";
import LinkCreate from './_link_create'

class PageUsersList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            rpl: {},
            load: false
        };

        this.responseServer = this.responseServer.bind(this);
    }

    responseServer(rpl) {
        this.setState({
            rpl: rpl.rpl,
            error: rpl.error,
            load: rpl.error === null
        });
    }

    componentDidMount() {
        RequestServer(window.location.pathname, 'post', null, this.responseServer)
    }

    render() {
        return (
            <div className='page-users-list'>
                <h1>Список пользователей</h1>
                <ul className='users-list'>
                    {this.renderUsersList()}
                </ul>

                <LinkCreate/>
            </div>
        )
    }

    renderUsersList() {
        switch (true) {
            case this.state.error:
                return <div className={'error'}>Error: {this.state.error.message}</div>;

            case this.state.rpl.errno > 0:
                return <div className={'error'}>Error: {this.state.rpl.error}</div>;

            case this.state.load:
                return (
                    <UserList usersList={this.state.rpl.list}/>
                )

        }
    }
}

class UserList extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            this.props.usersList.map(user => (
                <div key={user.id} className={'user-item'}>
                    <Link className={'wrapper-avatar'} style={{backgroundImage: "url(" + user.avatar + ")"}}
                          to={'/user/' + user.id}>
                    </Link>
                    <div className={'box-info'}>
                        <Link className={'wrapper-name hover-border'} to={'/user/' + user.id}>
                            {user.name}
                        </Link>
                        <div className="wrapper-created-at">
                            Создан:<br/>
                            <span dangerouslySetInnerHTML={{__html: formatDate(user.createdAt)}}/>
                        </div>
                    </div>
                </div>
            ))
        )
    }
}

export default PageUsersList;

function formatDate(date) {
    return date.replace(/[Z]/g, '').replace(/[T]/g, '<br/>в ').split('.')[0];
}