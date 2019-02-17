import React from "react"
import imgTrash from '../../../../images/trash.png'
import RequestServer from "../../../lib/request_server";

class BoxDelete extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            rpl: {},
            error: null,

            modal: false,
            name: props.name
        };

        this.id = props.id;
        this.onChangeFormStatus = props.onChangeFormStatus;

        this.onOpenModal = this.onOpenModal.bind(this);
        this.onCloseModal = this.onCloseModal.bind(this);
        this.onDelete = this.onDelete.bind(this);
        this.responseServer = this.responseServer.bind(this);
    }

    onOpenModal() {
        if (!this.state.modal) {
            this.setState({modal: true})
        }
    }

    onCloseModal() {
        this.setState({modal: false})
    }

    responseServer(rpl) {
        this.setState({
            rpl: rpl.rpl,
            error: rpl.error,
        });

        this.onChangeFormStatus(rpl.rpl.formStatus, '/users');
    }

    onDelete(e) {
        e.preventDefault();
        RequestServer('/user/delete/' + this.props.id, 'post', null, this.responseServer);
    }

    render() {
        switch (true) {
            case this.state.error:
                return <div className={'error'}>Error: {this.state.error.message}</div>;

            case this.state.rpl.errno > 0:
                return <div className={'error'}>Error: {this.state.rpl.error}</div>;

            default:
                return (
                    <div className={'box-delete ' + this.state.modal}
                         onClick={this.onOpenModal}
                         style={{backgroundImage: "url(../WWW/static/dist/" + imgTrash + ")"}}>

                        <div className={'content-modal ' + this.state.modal}>
                            <div className="text">Удалить?</div>
                            <div className="part-down">
                                <a href={'/user/delete/' + this.id} className={'modal-but yes'} onClick={this.onDelete}>Да</a>
                                <div className="modal-but no" onClick={this.onCloseModal}>Нет</div>
                            </div>
                        </div>
                    </div>
                )
        }
    }
}

export default BoxDelete