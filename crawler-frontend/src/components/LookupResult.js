import React, { Component }  from 'react';
import PropTypes from 'prop-types';
import Card, { CardContent } from 'material-ui/Card';
import Typography from 'material-ui/Typography';
import { withStyles } from 'material-ui/styles';

const styles = theme => ({
    card: {
        // paddingLeft: 0,
        // paddingBottom: 0,
        // marginBottom: 10,
        width: "100%",
        display: 'flex',
        // float: "left"
    },
});

class LookupResult extends Component {
    render(){
        const { classes } = this.props;
        return (
            <div>
                <Card className={classes.card}>
                    <CardContent>
                        <Typography type="headline" component="h2">
                            {this.props.details.name}
                        </Typography>
                        <Typography component="p">
                            {this.props.details.address}
                        </Typography>
                    </CardContent>
                </Card>
            </div>
        );
    }
}

LookupResult.propTypes = {
    classes: PropTypes.object.isRequired,
};
  
  export default withStyles(styles)(LookupResult);