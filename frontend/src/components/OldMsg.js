import React, {Component} from 'react';

class OldMsg extends Component {
    render() {
        const {content, speakerUserId} = this.props.msg;

        return (
            <div className="msgsInWindow"><strong>User {speakerUserId}</strong>: {content}</div>
        )
    }
}

export default OldMsg;
