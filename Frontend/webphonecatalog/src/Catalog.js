import React from 'react';

import axios from 'axios';

export default class Catalog extends React.Component {
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
            // <ul>
            //     { this.state.contacts.map(item =>
            //         <li key={item.id}>{item.name}</li>
            //         )}
            // </ul>
            <table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Phone</th>
                    </tr>
                </thead>
                <tfoot>
                    {this.state.contacts.map((item =>
                        <tr>
                            <td>{item.name}</td>
                            <td>{item.phone}</td>
                        </tr>
                    ))}
                </tfoot>
            </table>
        )}
}