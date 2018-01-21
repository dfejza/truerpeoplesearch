import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import PhoneSearchBar from "./../components/SearchBar";
import LookupResult from "./../components/LookupResult";

import axios from 'axios'; //HTTP client

const styles = theme => ({
    root: {
      flexGrow: 1,
      marginTop: "6%",
      marginLeft: "35%",
      marginRight: "35%",
      backgroundColor: "#f9fbfe",
    },
    headline:{
        paddingLeft: "16%",
        display: "inline",
        color: "#436280",
        fontWeight: "bold"
    },
    headline2:{
        display: "inline",
        color: "#a3bad2",
    },
    searchHeadline:{
        marginTop: 50,
        fontWeight: "bold",
        color: "#112236",
    }
});

// Point API requests to the right direction by modifying axios config
const instance = axios.create({
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    }
});

class Home extends Component {
    constructor(props) {
        super(props);
        this.state = {
            status: -5,
            result: {}
        };
        this.lookupNumber = this.lookupNumber.bind(this);
    }
    

    //HTTP call with user provided number
    lookupNumber(phoneNum, name){
        console.log(phoneNum);
        console.log(name);
        
        instance.post('/lookup', {
            number: phoneNum,
            name: (name === "" ? " " : name) // server doesnt play nice with empty strings TODO
          })
          .then(response => {
            console.log(response);
            this.setState({status : response.data.message.Ranking});
            // If there is a proper response, store it in the state
            if (response.data.message.Ranking){
                this.setState({result : {name : response.data.message.Name, address : response.data.message.Address}})
            }
          })
          .catch(error => {
            console.log(error);
            this.setState({status : -4})
        });
    }

    render() {
        const { classes } = this.props;
        return (
            <div className={classes.root}>
                <span>
                    <Typography type="display3" gutterBottom align="left" className={classes.headline}>
                        Truer
                    </Typography>
                    <Typography type="display3" gutterBottom align="left" className={classes.headline2}>
                        PeopleSearch
                    </Typography>
                </span>
                <Typography type="subheading" align="left" className={classes.searchHeadline}>
                    Reverse Phone Search
                </Typography>
                <PhoneSearchBar lookupNumber={this.lookupNumber} response={this.state.status}/>
                <Typography type="subheading" align="left" className={classes.searchHeadline}>
                    *Display purpose only. All requests will be blocked, requiring a captcha on the backed (AWS)
                </Typography>
                {this.state.status > -2 &&
                    <LookupResult details={this.state.result}/>
                }
            </div>
        );
    }
}

Home.propTypes = {
    classes: PropTypes.object.isRequired,
};

 export default withStyles(styles)(Home);