import React, {Component} from 'react';

class Friend extends Component {
    render() {

        const {name, id, nameNow, badgeNum,addChat} = this.props;
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
                    : null}
                    {badgeNum>0 && <span className="badgeNum">{badgeNum}</span>}
                    </div>
    }
}

export default Friend;
