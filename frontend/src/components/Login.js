import React, {Component} from 'react';

class Login extends Component {
    state = {
        name: "",
        showWarning: false
    }

    _change = (ev) => {
        const obj = {};
        const _name = ev.target.name;
        obj[_name] = ev.target.value;

        this.setState(obj);
    }

    _submit = (ev) => {
        ev.preventDefault();
        const {name} = {
            ...this.state
        };

        if (name) {
            this
                .props
                .signup({name});
        } else {
            this.setState({showWarning: true})
        }
    }

    render() {
        const {showWarning} = this.state;
        return (
            <div>
                <form onSubmit={this._submit}>
                    <input name="name" placeholder="Name" onChange={this._change}/> {showWarning && <div>Please use a valid name</div>}

                    <input type="submit" value="Sign me up"/>
                </form>
            </div>
        );
    }
}

export default Login;
