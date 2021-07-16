import React from 'react';

import axios from 'axios';

export default class PersonList extends React.Component {
    state = {
        contacts: []
    }

    componentDidMount() {
        axios.get(`http://contacts.com.br:3001/contacts`)
            .then(res => {
                const contacts = res.data;
                this.setState({ contacts });
            })
    }

    render() {
        return (
            <ul>
                { this.state.contacts.map(contact => <li>{contact.name}</li>)}
            </ul>
        )
    }
}