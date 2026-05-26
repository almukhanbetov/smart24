package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUsers(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT id, phone, fullname, created_at, user_code, avatar FROM users ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetDevices(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT * FROM devices ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetDeviceSet(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT * FROM device_set ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetTariffs(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT * FROM tariffs ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetPayments(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT * FROM payments ORDER BY id DESC LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetCoin(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT * FROM coin ORDER BY id DESC LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetMoney(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset := LimitOffset(c)
		QueryAll(c, pool, `SELECT * FROM money ORDER BY id DESC LIMIT $1 OFFSET $2`, limit, offset)
	}
}

func GetDeviceFull(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		account := c.Param("account")
		QueryAll(c, pool, `
			SELECT
				d.id,
				d.account,
				d.code_device,
				d.user_code,
				d.device_name,
				d.type,
				d.bin,
				d.gruppa,
				ds.user_wifi,
				ds.signal_wifi,
				ds.status,
				ds.data_status,
				ds.l,
				ds.n,
				ds.data_inkas
			FROM devices d
			LEFT JOIN device_set ds ON ds.account = d.account
			WHERE d.account = $1
		`, account)
	}
}

func GetDashboard(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		var usersCount, devicesCount, paymentsCount int
		var paymentsSum float64

		_ = pool.QueryRow(ctx, `SELECT COUNT(*) FROM users`).Scan(&usersCount)
		_ = pool.QueryRow(ctx, `SELECT COUNT(*) FROM devices`).Scan(&devicesCount)
		_ = pool.QueryRow(ctx, `SELECT COUNT(*) FROM payments`).Scan(&paymentsCount)
		_ = pool.QueryRow(ctx, `SELECT COALESCE(SUM(sum), 0) FROM payments`).Scan(&paymentsSum)

		c.JSON(http.StatusOK, gin.H{
			"users_count":    usersCount,
			"devices_count":  devicesCount,
			"payments_count": paymentsCount,
			"payments_sum":   paymentsSum,
		})
	}
}

func GetMaxData(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		QueryAll(c, pool, `
			SELECT
				d.id,
				d.account,
				d.code_device,
				d.device_name,
				d.type,
				d.bin,
				d.gruppa,
				u.id AS user_id,
				u.phone,
				u.fullname,
				u.user_code,
				u.avatar,
				ds.user_wifi,
				ds.signal_wifi,
				ds.status,
				ds.data_status,
				ds.l,
				ds.n,
				ds.data_inkas,
				COALESCE(p.total_payments, 0) AS total_payments,
				COALESCE(p.payments_count, 0) AS payments_count,
				COALESCE(m.total_money, 0) AS total_money,
				COALESCE(cn.total_coin, 0) AS total_coin
			FROM devices d
			LEFT JOIN users u ON u.user_code = d.user_code
			LEFT JOIN device_set ds ON ds.account = d.account
			LEFT JOIN (
				SELECT account, SUM(sum) AS total_payments, COUNT(*) AS payments_count
				FROM payments
				GROUP BY account
			) p ON p.account = d.account
			LEFT JOIN (
				SELECT account::text AS account, SUM(pay_money) AS total_money
				FROM money
				GROUP BY account
			) m ON m.account = d.account
			LEFT JOIN (
				SELECT account::text AS account, SUM(pay_coin) AS total_coin
				FROM coin
				GROUP BY account
			) cn ON cn.account = d.account
			ORDER BY d.id
		`)
	}
}
