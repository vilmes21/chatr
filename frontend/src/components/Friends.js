import React, {Component} from 'react';
import Friend from './Friend';

class Friends extends Component {
    render() {
        const {friends, nameNow,addChat} = this.props;

        const nameArr = Object.keys(friends);
        if (nameArr.length === 0){
            return <span>No friend yet. Poor lonely you.</span>
        }
        
        return nameArr.map(k => <Friend addChat={addChat} key={friends[k]} name={k} id={friends[k]} nameNow={nameNow}/>)
    }
}

export default Friends;
