import React, {Component} from 'react';
import OldMsg from './OldMsg';

class OldMsgs extends Component {
    render() {
        const {oldMsgs} = this.props;

        if (oldMsgs.length === 0){
            return null
        }
        
        return oldMsgs.map((msg, index) => <OldMsg key={index} msg={msg}/>)
    }
}

export default OldMsgs;
