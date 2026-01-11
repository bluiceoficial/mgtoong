// Copyright (C) 2025-2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package mgtoong

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type MGTOONG struct {
	sType       string
	sFields     []string
	sRecords    []map[string]string
	sPrimaryKey string
}

func (m *MGTOONG) LoadFile(filename string, primaryKey string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return m.LoadTOON(string(data), primaryKey)
}

func (m *MGTOONG) LoadTOON(text string, primaryKey string) error {
	toon, err := parse(text, primaryKey)
	if err != nil {
		return err
	}

	m.sType = toon.sType
	m.sFields = toon.sFields
	m.sRecords = toon.sRecords
	m.sPrimaryKey = toon.sPrimaryKey
	return nil
}

func (m *MGTOONG) Create(sType string, fields []string, primaryKey string) error {
	m.sType = sType
	m.sFields = fields
	m.sPrimaryKey = primaryKey

	for _, f := range fields {
		if f == primaryKey {
			return nil
		}
	}

	return fmt.Errorf("campo primário '%s' não existe nos campos definidos", primaryKey)
}

func parse(text string, primaryKey string) (*MGTOONG, error) {
	lines := []string{}
	for _, l := range strings.Split(text, "\n") {
		l = strings.TrimSpace(l)
		if l != "" {
			lines = append(lines, l)
		}
	}

	if len(lines) == 0 {
		return nil, errors.New("TOON vazio")
	}

	re := regexp.MustCompile(`^([a-zA-Z0-9_]+)\[([^\]]+)\]$`)
	matches := re.FindStringSubmatch(lines[0])
	if matches == nil {
		return nil, errors.New("cabeçalho inválido em TOON")
	}

	typ := matches[1]
	fields := strings.Split(matches[2], "|")

	toon := &MGTOONG{}
	if err := toon.Create(typ, fields, primaryKey); err != nil {
		return nil, err
	}

	for i := 1; i < len(lines); i++ {
		values := strings.Split(lines[i], "|")
		if len(values) != len(fields) {
			return nil, fmt.Errorf("linha %d inválida: número de colunas não corresponde ao cabeçalho", i)
		}

		record := map[string]string{}
		for j, f := range fields {
			record[f] = strings.TrimSpace(values[j])
		}
		toon.sRecords = append(toon.sRecords, record)
	}

	return toon, nil
}

func (m *MGTOONG) ToString() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("%s[%s]", m.sType, strings.Join(m.sFields, "|")))

	for _, r := range m.sRecords {
		values := []string{}
		for _, f := range m.sFields {
			values = append(values, r[f])
		}
		lines = append(lines, strings.Join(values, "|"))
	}

	return strings.Join(lines, "\n")
}

func (m *MGTOONG) Add(record map[string]string) (map[string]string, error) {
	key := m.sPrimaryKey
	val := record[key]

	if val == "" {
		return nil, fmt.Errorf("campo primário '%s' é obrigatório", key)
	}

	if m.exists(val) {
		return nil, fmt.Errorf("registro com %s='%s' já existe", key, val)
	}

	filtered := map[string]string{}
	for _, f := range m.sFields {
		filtered[f] = record[f]
	}

	m.sRecords = append(m.sRecords, filtered)
	return filtered, nil
}

func (m *MGTOONG) ReadOne(keyValue string) map[string]string {
	key := m.sPrimaryKey
	for _, r := range m.sRecords {
		if r[key] == keyValue {
			return r
		}
	}
	return nil
}

func (m *MGTOONG) ReadAll() []map[string]string {
	return m.sRecords
}

func (m *MGTOONG) Update(keyValue string, newData map[string]string) bool {
	key := m.sPrimaryKey

	for _, r := range m.sRecords {
		if r[key] == keyValue {
			for k, v := range newData {
				if contains(m.sFields, k) {
					r[k] = v
				}
			}
			return true
		}
	}
	return false
}

func (m *MGTOONG) Delete(keyValue string) bool {
	key := m.sPrimaryKey

	for i, r := range m.sRecords {
		if r[key] == keyValue {
			m.sRecords = append(m.sRecords[:i], m.sRecords[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MGTOONG) exists(keyValue string) bool {
	key := m.sPrimaryKey
	for _, r := range m.sRecords {
		if r[key] == keyValue {
			return true
		}
	}
	return false
}

func contains(arr []string, v string) bool {
	for _, s := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func (m *MGTOONG) SaveFile(filename string) error {
	return os.WriteFile(filename, []byte(m.ToString()), 0644)
}

func Validate(text string, primaryKey string) map[string]interface{} {
	toon, err := parse(text, primaryKey)
	if err != nil {
		return map[string]interface{}{
			"valid": false,
			"error": err.Error(),
		}
	}

	pk := toon.sPrimaryKey
	seen := map[string]bool{}

	for _, r := range toon.sRecords {
		v := r[pk]
		if seen[v] {
			return map[string]interface{}{
				"valid": false,
				"error": fmt.Sprintf("valores duplicados no campo primário '%s'", pk),
			}
		}
		seen[v] = true
	}

	return map[string]interface{}{
		"valid":      true,
		"type":       toon.sType,
		"primaryKey": pk,
		"fields":     toon.sFields,
		"records":    len(toon.sRecords),
	}
}
