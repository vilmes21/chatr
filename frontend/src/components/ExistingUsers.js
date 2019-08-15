import React, {Component} from 'react';

class ExistingUsers extends Component {
    render() {
        const {users, login} = this.props;

        return Object
            .keys(users)
            .map(x => {
                return <div
                    onClick={() => {
                    login({name: x, id: users[x]})
                }}>{users[x]}. {x}</div>
            })
    }
}

export default ExistingUsers;
