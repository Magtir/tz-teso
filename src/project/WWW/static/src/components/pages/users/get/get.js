import React from "react";
import RequestServer from '../../../lib/request_server'
import "./get.scss";
import BoxName from "./_box_name"
import BoxAvatar from "./_box_avatar"
import BoxDelete from "./_box_delete"
import {Redirect} from 'react-router-dom'
import Success from '../../../_parts/_success'
import LinkBack from '../../../_parts/_link_back'

class PageUser extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            rpl: {},
            load: false,
            formStatus: false,
            redirect: ''
        };

        this.timer = null;

        this.onChangeFormStatus = this.onChangeFormStatus.bind(this);
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

    onChangeFormStatus(status, urlRedirect = '') {
        clearTimeout(this.timer);

        this.setState({formStatus: status});

        if (status) {
            this.timer = setTimeout(() => {
                this.setState({formStatus: false, redirect: urlRedirect});
            }, 1000);
        }
    }

    render() {
        return (
            <div className='page-user'>
                <LinkBack/>
                <Success status={this.state.formStatus}/>
                {this.renderUser()}
            </div>
        )
    }

    renderUser() {
        switch (true) {
            case this.state.error:
                return <div className={'error'}>Error: {this.state.error.message}</div>;

            case this.state.rpl.errno > 0:
                return <div className={'error'}>Error: {this.state.rpl.error}</div>;

            case this.state.redirect !== '':
                return <Redirect to={this.state.redirect}/>;

            case this.state.load:
                let user = this.state.rpl.user;

                return (
                    <div className='block-user'>
                        <h1>Пользователь</h1>
                        <BoxName
                            id={user.id}
                            name={user.name}
                            onChangeFormStatus={this.onChangeFormStatus}
                        />
                        <BoxAvatar
                            avatar={user.avatar}
                        />

                        <div className="box-created-at">
                            Создан {formatDate(user.createdAt)}
                        </div>

                        <BoxDelete
                            id={user.id}
                            name={user.name}
                            onChangeFormStatus={this.onChangeFormStatus}
                        />
                    </div>
                )

        }
    }
}

function formatDate(date) {
    return date.replace(/[Z]/g, '').replace(/[T]/g, ' в ').split('.')[0];
}

export default PageUser;