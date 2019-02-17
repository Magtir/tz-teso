import React from "react";
import imgEdit from '../../../../images/edit.png'
import imgSave from '../../../../images/save.png'
import RequestServer from "../../../lib/request_server";

class BoxName extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            rpl: {},
            error: null,

            edit: false,
            name: props.name
        };

        this.id = props.id;
        this.onChangeFormStatus = props.onChangeFormStatus;

        this.onEdit = this.onEdit.bind(this);
        this.onNameChange = this.onNameChange.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.responseServer = this.responseServer.bind(this);
    }

    onEdit() {
        this.setState((prevState) => ({
            edit: !prevState.edit
        }));
    }

    onNameChange(e) {
        this.setState({name: e.target.value});
    }


    responseServer(rpl) {
        this.setState({
            rpl: rpl.rpl,
            error: rpl.error,
        });

        this.onChangeFormStatus(rpl.rpl.formStatus);

        if (rpl.rpl.formStatus) {
            this.onEdit();
        }
    }

    onSubmit(e) {
        e.preventDefault();
        RequestServer('/user/edit/' + this.props.id, 'post', this.state.name, this.responseServer);
    }

    render() {
        let name = this.state.name;

        switch (true) {
            case this.state.error:
                return <div className={'error'}>Error: {this.state.error.message}</div>;

            case this.state.rpl.errno > 0:
                return <div className={'error'}>Error: {this.state.rpl.error}</div>;

            case !this.state.edit:
                return (
                    <div className={'box-name'}>
                        <div className={'name'}>
                            {name}
                        </div>
                        <div className={'icon-edit edit'}
                             onClick={this.onEdit}
                             style={{backgroundImage: "url(../WWW/static/dist/" + imgEdit + ")"}}>
                        </div>
                    </div>
                );

            case this.state.edit:
                return (
                    <div className={'box-name'}>
                        <form onSubmit={this.onSubmit}>
                            <div className={'field-error'}>
                                {this.state.rpl.formError}
                            </div>
                            <label>
                                <input
                                    autoFocus
                                    className={'input-name'}
                                    type="text"
                                    name="name"
                                    value={name}
                                    onChange={this.onNameChange}
                                />
                            </label>
                            <input className={'icon-edit save'}
                                   type="submit"
                                   value={''}
                                   style={{backgroundImage: "url(../WWW/static/dist/" + imgSave + ")"}}/>
                        </form>
                    </div>
                )
        }
    }
}

export default BoxName;