import React from "react";
import RequestServer from '../../../lib/request_server'
import "./create.scss";
import { Redirect } from 'react-router-dom'
import LinkBack from '../../../_parts/_link_back'
import imgCreate from '../../../../images/create-w.png'

class PageUserCreate extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            rpl: {},
            error: null,

            name: '',
            avatar: '',
            id: ''
        };

        this.onNameChange = this.onNameChange.bind(this);
        this.onAvatarChange = this.onAvatarChange.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.responseServer = this.responseServer.bind(this);
    }

    onNameChange(e) {
        this.setState({name: e.target.value});
    }

    onAvatarChange(e) {
        this.setState({avatar: e.target.value});
    }

    responseServer(rpl) {
        this.setState({
            rpl: rpl.rpl,
            error: rpl.error,
            id: rpl.rpl.id
        })
    }

    onSubmit(e) {
        e.preventDefault();

        let body = JSON.stringify({name: this.state.name, avatar: this.state.avatar});
        RequestServer(window.location.pathname, 'post', body, this.responseServer);
    }

    render() {
        return (
            <div className='page-users-create'>
                <LinkBack />
                <h1>Создание пользователя</h1>
                    {this.renderUserCreate()}
            </div>
        )
    }

    renderUserCreate() {
        switch (true) {
            case this.state.error:
                return <div className={'error'}>Error: {this.state.error.message}</div>;

            case this.state.rpl.errno > 0:
                return <div className={'error'}>Error: {this.state.rpl.error}</div>;

            case this.state.id > 0:
                return <Redirect to={'/user/' + this.state.id} />;

            default:
                let name = this.state.name;
                let avatar = this.state.avatar;

                return (
                    <form onSubmit={this.onSubmit}>

                        <div className="field">
                            <div className={'field-error'}>
                                {this.state.rpl.formNameError}
                            </div>
                            <label> Name:
                                <input
                                    autoFocus
                                    className={'input-name'}
                                    type="text"
                                    name="name"
                                    value={name}
                                    onChange={this.onNameChange}
                                />
                            </label>
                        </div>

                        <div className="field">
                            <div className={'field-error'}>
                                {this.state.rpl.formAvatarError}
                            </div>
                            <label> Avatar:
                                <input
                                    className={'input-avatar'}
                                    type="text"
                                    name="avatar"
                                    value={avatar}
                                    onChange={this.onAvatarChange}
                                />
                            </label>
                        </div>

                        <input className={'icon-create'}
                               type="submit"
                               value={''}
                               style={{backgroundImage: "url(../WWW/static/dist/" + imgCreate + ")"}}
                        />
                    </form>
                )
        }
    }
}

export default PageUserCreate;