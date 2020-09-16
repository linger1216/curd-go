package svc

/**
 * @api {post} /eg/echo Create Echo
 * @apiName CreateEcho
 * @apiGroup Echo
 *
 * @apiParam {Object[]} echos  New echo's id.
 * @apiParam {String} echos.id  New echo's id.
 * @apiParam {String} echos.age  New echo's age.
 * @apiParam {String} echos.name  New echo's name.
 * @apiParam {String} echos.geometry  标准的 geometry json.
 * @apiParam {String[]} echos.books  New echo's books.
 * @apiParam {String[]} echos.tags  New echo's tags.
 *
 *
 * @apiParamExample {json} Request-Example:
 *	[
 *		{
 *			"age": 18,
 *			"name": "lid",
 *			"geometry": {
 *				"type": "Point",
 *				"coordinates": [
 *					39.375,
 *					58.26328705248601
 *				]
 *			},
 *			"books": [
 *				"book1",
 *				"book2"
 *			],
 *      "tags": [
 *				1,
 *				2
 *			]
 *		}
 *	]
 *
 *
 * @apiSuccess {String[]} id list  添加后的id列表
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 * [
 *   "1306122237937455104"
 * ]
 *
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 internal server error
 *     HTTP/1.1 404 not found
 *     HTTP/1.1 400 invalid para
 */

/**
 * @api {delete} /eg/echo/:ids Delete Echo
 * @apiName DeleteEcho
 * @apiGroup Echo
 *
 * @apiParam {String} ids  echo's id, 多个以逗号分割
 *
 *
 * @apiParamExample {json} Request-Example:
 *	/eg/echo/1306122237937455104,1306122237937455105
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 internal server error
 *     HTTP/1.1 404 not found
 *     HTTP/1.1 400 invalid para
 */

/**
 * @api {put} /eg/echo Update Echo
 * @apiName UpdateEcho
 * @apiGroup Echo
 *
 * @apiParam {Object[]} echos  New echo's id.
 * @apiParam {String} echos.id  New echo's id.
 * @apiParam {String} echos.age  New echo's age.
 * @apiParam {String} echos.name  New echo's name.
 * @apiParam {String} echos.geometry  标准的 geometry json.
 * @apiParam {String[]} echos.books  New echo's books.
 * @apiParam {String[]} echos.tags  New echo's tags.
 *
 *
 * @apiParamExample {json} Request-Example:
 *	[
 *		{
 *			"age": 18,
 *			"name": "lid",
 *			"geometry": {
 *				"type": "Point",
 *				"coordinates": [
 *					39.375,
 *					58.26328705248601
 *				]
 *			},
 *			"books": [
 *				"book1",
 *				"book2"
 *			],
 *      "tags": [
 *				1,
 *				2
 *			]
 *		}
 *	]
 *
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 internal server error
 *     HTTP/1.1 404 not found
 *     HTTP/1.1 400 invalid para
 */

/**
 * @api {head} /eg/echo List Echo Count
 * @apiName ListEchoCount
 * @apiGroup Echo
 *
 * @apiParam {String} age  age.
 * @apiParam {String} name  name.
 * @apiParam {String} book  book.
 * @apiParam {String} tag  tag.
 * @apiParam {Number} longitude  longitude.
 * @apiParam {Number} latitude   latitude.
 * @apiParam {String} spatial_reference [spatialReference=gcj02]  坐标系.
 * @apiParam {Number} radius  [radius=1000] 搜索半径.
 * @apiParam {Number} start_time  开始时间.
 * @apiParam {String} end_time  [end_time=当前时间] 结束时间.
 *
 *
 * @apiParamExample {json} Request-Example:
 *
 * name=lid&age=29&book=book1&tag=1&longitude=39.375&latitude=58.263287&radius=10&current_page=0&page_size=1&start_time=1600073795&end_time=1600073797&spatial_reference=gcj02
 *
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     X-Total-Count: 1
 *
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 internal server error
 *     HTTP/1.1 404 not found
 *     HTTP/1.1 400 invalid para
 */

/**
 * @api {get} /eg/echo List Echo
 * @apiName ListEcho
 * @apiGroup Echo
 *
 * @apiParam {String} age  age.
 * @apiParam {String} name  name.
 * @apiParam {String} book  book.
 * @apiParam {String} tag  tag.
 * @apiParam {Number} longitude  longitude.
 * @apiParam {Number} latitude   latitude.
 * @apiParam {String} spatial_reference [spatialReference=gcj02]  坐标系.
 * @apiParam {Number} radius  [radius=1000] 搜索半径.
 * @apiParam {Number} start_time  开始时间.
 * @apiParam {String} end_time  [end_time=当前时间] 结束时间.
 * @apiParam {Number} current_page  [current_page=0] 当前页.
 * @apiParam {Number} page_size  [page_size=10] 页大小.
 *
 *
 * @apiParamExample {json} Request-Example:
 *
 * name=lid&age=29&book=book1&tag=1&longitude=39.375&latitude=58.263287&radius=10&current_page=0&page_size=1&start_time=1600073795&end_time=1600073797&spatial_reference=gcj02
 *
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *	[
 *		{
 *			"age": 18,
 *			"name": "lid",
 *			"geometry": {
 *				"type": "Point",
 *				"coordinates": [
 *					39.375,
 *					58.26328705248601
 *				]
 *			},
 *			"books": [
 *				"book1",
 *				"book2"
 *			],
 *      "tags": [
 *				1,
 *				2
 *			]
 *		}
 *	]
 *
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 internal server error
 *     HTTP/1.1 404 not found
 *     HTTP/1.1 400 invalid para
 */

/**
 * @api {get} /eg/echo/:ids Get Echo
 * @apiName GetEcho
 * @apiGroup Echo
 *
 * @apiParam {String} ids  echo's id, 多个以逗号分割
 *
 *
 * @apiParamExample {json} Request-Example:
 *	/eg/echo/1306122237937455104,1306122237937455105
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *	[
 *		{
 *			"age": 18,
 *			"name": "lid",
 *			"geometry": {
 *				"type": "Point",
 *				"coordinates": [
 *					39.375,
 *					58.26328705248601
 *				]
 *			},
 *			"books": [
 *				"book1",
 *				"book2"
 *			],
 *      "tags": [
 *				1,
 *				2
 *			]
 *		}
 *	]
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 internal server error
 *     HTTP/1.1 404 not found
 *     HTTP/1.1 400 invalid para
 */
