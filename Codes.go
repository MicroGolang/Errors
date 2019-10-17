/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Thursday 26 September 2019 - 11:04:11
** @Filename:				codes.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Thursday 17 October 2019 - 15:20:07
*******************************************************************************/

package		errors

const CLIENT_400 = `Something was wrong with the request`
const CLIENT_401 = `You are not allowed to access this ressource`
const CLIENT_403 = `You are not allowed to access this ressource`
const CLIENT_404 = `One or many elements were not found`
const CLIENT_409 = `You are facing a conflict. Please verify what you want and the existing data`
const CLIENT_440 = `Login informations timeout`
const CLIENT_498 = `Your token is invalid`
const CLIENT_499 = `You need a token to access this ressource`

const SERV_500 = `Unexpected error`