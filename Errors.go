/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Thursday 26 September 2019 - 11:04:11
** @Filename:				error.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Thursday 26 September 2019 - 12:07:06
*******************************************************************************/

package		error

import		"errors"
import		"net/http"
import		"gitlab.com/gopackages/logs"
import		"gitlab.com/gopackages/restresponse"

/*Handle **********************************************************************
*	Handler an error with a recover in order to stop the current process and
*	to resolve an error to the client
******************************************************************************/
func	Handle(w http.ResponseWriter) {
	if err := recover(); (err != nil) {
		response.ResolveWithError(w, New(`Unexpected Error`), 500)
	}
}
/*HandleSimple ****************************************************************
*	Handler an error with a recover in order to stop the current process
******************************************************************************/
func	HandleSimple() {
	if err := recover(); (err != nil) {
		logs.Error(`Recovering`)
		logs.Pretty(err)
		return
	}
}

/*HasError ********************************************************************
*	Log l'erreur et return true en cas d'erreur
******************************************************************************/
func	HasError(err error) bool {
	if (err != nil) {
		logs.ErrorCrash(err.Error())
		return true
	}
	return false
}

/*HasStrError *****************************************************************
*	Log l'erreur (en cas de string) et return true en cas d'erreur
******************************************************************************/
func	HasStrError(err string) bool {
	if (err != ``) {
		logs.ErrorCrash(err)
		return true
	}
	return false
}

/*New *************************************************************************
*	Shorthand autour de errors.new pour n'importer qu'un seul package
******************************************************************************/
func	New(err string) error {
	return errors.New(err)
}
