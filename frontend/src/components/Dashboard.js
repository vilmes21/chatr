import React, {Component} from 'react';
import fakeFriends from '../fakeData/fakeFriends'
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    withRouter
  } from 'react-router-dom'
import Friends from './Friends'
import ChatWindows from './ChatWindows'


class Dashboard extends Component {
    state = {
        friends: false,
        chatingWith: []
    }

    componentDidMount=()=>{
        //Demo fake ajax
        setTimeout(()=>{
            this.setState({friends: fakeFriends})
        }, 0)
    }

    addChat = userObj => {
        const {chatingWith}=this.state;
        const alreadyInChat = chatingWith.find(x => x.id === userObj.id);
        if (alreadyInChat) return;
        
        this.setState({
            chatingWith: [...chatingWith, userObj]
        })
    }
     
    render() {
        const {isAuth,userNowId, nameNow} = this.props;

        if (!isAuth){
            return <Redirect to='/login' />
        }

        const {friends, chatingWith}=this.state;
        
        return (
            <div>
               <h3>
               Friends to Chat with:
               </h3>

               {friends? <Friends addChat={this.addChat} nameNow={nameNow} friends={friends}/>: <span>loading...</span>}

               <ChatWindows userNowId={userNowId} users={chatingWith}/>
            </div>
        );
    }
}

export default Dashboard;
