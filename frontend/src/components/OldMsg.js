import React, {Component} from 'react';

class OldMsg extends Component {
    render() {
        const {content, chatSpeakerId} = this.props.msg;

        return (
            <div><strong>Speaker {chatSpeakerId}</strong>: {content}</div>
        )
    }
}

export default OldMsg;
