import React, {Component} from 'react';

class Friend extends Component {
    render() {

        const {name, id, nameNow, addChat} = this.props;
        const isSelf = name === nameNow;
        const _onclick = isSelf
            ? () => {
                alert("Self-talk is bad for the soul")
            }
            : () => {
                addChat({name, id})
            }

            return <div onClick={_onclick}>{id}. {name}
                {isSelf
                    ? "(you)"
                    : null}</div>
    }
}

export default Friend;
