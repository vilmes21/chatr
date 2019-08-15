import React, {Component} from 'react';

class Friend extends Component {
    render() {
    
        const {name,id,nameNow, addChat} = this.props;

        return <div onClick={() =>{
            addChat({
                name,
                id
            })
        }}>{id}. {name} {name===nameNow? "(you)":null}</div>
    }
}

export default Friend;
