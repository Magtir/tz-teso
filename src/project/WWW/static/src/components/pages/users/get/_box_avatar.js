import React from "react";

class BoxAvatar extends React.Component {
    constructor(props) {
        super(props);

        this.avatar = props.avatar;
    }

    render() {
        return (
            <div className={'box-avatar'} style={{backgroundImage: "url(" + this.avatar + ")"}}>
            </div>
        )
    }
}

export default BoxAvatar